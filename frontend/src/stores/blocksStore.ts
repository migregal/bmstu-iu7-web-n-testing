import { ApiCallMethods } from "hooks/apiCall";
import { AuthStore } from "./authStore";


export interface Block {
  id: string
  blocked_until: string
}

export interface BlocksStore {
  get(params: { id: string }): Promise<Block>,
  block(params: { id: string, until: Date }): Promise<Block>,
  unblock(params: { id: string }): Promise<void>,
}

export function createBlocksStore(
  auth: AuthStore, api: ApiCallMethods<Block>
): BlocksStore {
  return {
    async get({ id }) {
      return api.one({
        resourceID: id,
        options: { auth: auth.user.token, resp: 'json' }
      })
    },
    async block({ id, until }) {
      return api.update({
        resourceID: id,
        payload: { "until": until },
        options: { auth: auth.user.token }
      })
    },
    async unblock({ id }) {
      return api.destroy({
        resourceID: id,
        options: { auth: auth.user.token }
      })
    }
  }
}
