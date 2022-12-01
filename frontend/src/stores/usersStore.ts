import { ApiCallMethods } from "hooks/apiCall";
import { runInAction } from "mobx";
import { AuthStore } from "./authStore";
import { Block, BlocksStore } from "./blocksStore";

const usersKey = "users"

export interface User {
  id: string
  username: string
  email: string
  fullname: string
  block: string
}

export interface UsersStore {
  users: { infos: User[], total: number },
  getAll(params: { page: number, perPage: number }): Promise<User[]>,
  get(params: { ids: string[] }): Promise<User[]>,
  block(params: { ids: string[], until: Date }): Promise<Block[]>,
  unblock(params: { ids: string[] }): Promise<void[]>,
}

export function createUsersStore(
  auth: AuthStore, blocks: BlocksStore, api: ApiCallMethods<User>
): UsersStore {
  return {
    users: { infos: [], total: 0 },
    async getAll({ page, perPage }) {
      const body = await api.all({
        page: page,
        perPage: perPage,
        options: { auth: auth.user.token, resp: 'json' }
      })

      if (auth.isadmin()) {
        const b = await Promise
          .all(body.infos.map((u) => blocks.get({ id: u.id })))

        body.infos = body.infos.map((i) => {
          i.block = b.find((el) => el.id == i.id)?.blocked_until ?? ''
          return i
        })
      }

      runInAction(() => {
        this.users.infos = body.infos
        this.users.total = body.total
      })

      localStorage.setItem(usersKey, JSON.stringify(this.users));
      return this.users.infos
    },
    async get({ ids }) {
      const body = await Promise
        .all(ids.map((id) => api.one({
          resourceID: id,
          options: { auth: auth.user.token, resp: 'json' }
        })))

      runInAction(() => {
        this.users.infos = body
        this.users.total = 0
      })

      localStorage.setItem(usersKey, JSON.stringify(this.users));
      return this.users.infos
    },
    async block({ ids, until }) {
      return Promise
        .all(ids.map((id) => blocks.block({ id: id, until: until })))
    },
    async unblock({ ids }) {
      return Promise
        .all(ids.map((id) => blocks.unblock({ id: id })))
    }
  }
}
