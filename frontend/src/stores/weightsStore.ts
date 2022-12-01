import { runInAction } from "mobx"

import { ApiCallMethods } from "hooks/apiCall"

import { AuthStore } from "./authStore"

export interface Weight {
  id: string
  name: string
}

export interface WeightsStore {
  weights: Weight[],
  getByStructureID(structureID: string): Promise<Weight[]>
  add(params: { modelID: string, title: string, weights: File }): Promise<void>
  update(params: { modelID: string, id: string, title: string, weights: File }): Promise<void>
}

export function createWeightsStore(auth: AuthStore, api: ApiCallMethods<Weight>): WeightsStore {
  return {
    weights: [],
    async getByStructureID(structureID) {
      const body = await api.all({
        query: { "structure_id": structureID },
        options: { auth: auth.user.token, resp: "json" },
      })

      runInAction(() => {
        this.weights = body.infos
      })

      return body.infos
    },
    async add({ modelID, title, weights }) {
      await api.create({
        payload: {
          'model_id': modelID,
          'weights_title': title,
          'weights': weights
        },
        options: { auth: auth.user.token }
      })

      return
    },
    async update({ modelID, id, title, weights }) {
      await api.update({
        resourceID: id,
        payload: {
          'model_id': modelID,
          'weights_title': title,
          'weights': weights
        },
        options: { auth: auth.user.token }
      })

      return
    }
  }
}
