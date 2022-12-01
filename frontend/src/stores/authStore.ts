import { runInAction } from "mobx"

import { isExpired, decodeToken } from "react-jwt"

import { ApiCallMethods } from "hooks/apiCall"

const authKey = "auth"

const userRole = 1 << 0
const adminRole = 1 << 1
const statRole = 1 << 2

export interface Auth {
  token: string | null,
}

const getUser = (): Auth => {
  const stored = localStorage.getItem(authKey)
  if (!stored)
    return {} as Auth

  return JSON.parse(stored) as Auth
}

export interface AuthStore {
  user: Auth,
  register(form: { username: string, email: string, fullname: string, password: string }): Promise<void>,
  authorize(login: string, password: string): Promise<void>,
  refresh(): void,
  unautorize(): void
  authorized(): boolean
  id(): string
  isadmin(): boolean
  isregular(): boolean
  isstat(): boolean
}

export const createAuthStore = (api: ApiCallMethods<Auth>): AuthStore => {
  return {
    user: getUser(),
    async register({ username, email, fullname, password }) {
      const body = await api.create({
        payload: {
          username: username,
          email: email,
          fullname: fullname,
          password: password,
        },
        options: { resp: "json", suffix: "registration" }
      })

      runInAction(() => {
        this.user.token = body.token
      })

      localStorage.setItem(authKey, JSON.stringify(this.user));
    },
    async authorize(login: string, password: string) {
      const body = await api.create({
        payload: { email: login, password: password },
        options: { resp: "json", suffix: "login" }
      })

      runInAction(() => {
        this.user.token = body.token
      })

      localStorage.setItem(authKey, JSON.stringify(this.user));
    },
    async refresh() {
      const body = await api.create({
        options: { resp: "json", suffix: "refresh"}
      })

      runInAction(() => {
        this.user.token = body.token
      })

      localStorage.setItem(authKey, JSON.stringify(this.user));
    },
    unautorize() {
      localStorage.removeItem(authKey);

      runInAction(() => {
        this.user.token = null
      })
    },
    authorized(): boolean {
      return !!this.user.token
        && !isExpired(this.user.token)
    },
    id(): string {
      const decodedToken = decodeToken(this.user.token ?? '')
      return (decodedToken as { user_id: string })?.user_id ?? null
    },
    isadmin(): boolean {
      if (!this.authorized())
        return false;

      const decodedToken = decodeToken(this.user.token ?? '') as { flags: number }
      return ((decodedToken?.flags ?? 0) & adminRole) !== 0
    },
    isregular(): boolean {
      if (!this.authorized())
        return false;

      const decodedToken = decodeToken(this.user.token ?? '') as { flags: number }
      return ((decodedToken?.flags ?? 0) & userRole) != 0
    },
    isstat(): boolean {
      if (!this.authorized())
        return false;

      const decodedToken = decodeToken(this.user.token ?? '') as { flags: number }
      return ((decodedToken?.flags ?? 0) & statRole) != 0
    }
  }
}
