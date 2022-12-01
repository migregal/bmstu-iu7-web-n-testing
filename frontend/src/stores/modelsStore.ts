import JSZip from "jszip";

import { ApiCallMethods } from "hooks/apiCall";
import { runInAction } from "mobx";
import { AuthStore } from "./authStore";
import { UsersStore } from "./usersStore";

const modelsKey = "models"

interface Structure {
  id: string
  layers: never[]
  links: never[]
  neurons: never[]
  weights?: { id: string }[]
}
export interface Model {
  id: string
  title: string
  owner_id: string
  owner_name?: string
  structure: Structure
}

export interface ModelsStore {
  models: { infos: Model[], total: number },
  getAll(params: { page: number, perPage: number }): Promise<Model[]>,
  get(params: { id: string }): Promise<Model>,
  download(params: { ids: string[] }): Promise<void>,
  upload(params: { title: string, structure: File, weights: File }): Promise<void>
  delete(params: { ids: string[] }): Promise<void[]>,
}

export function createModelsStore(
  auth: AuthStore, users: UsersStore, api: ApiCallMethods<Model>
): ModelsStore {
  return {
    models: { infos: [], total: 0 },
    async getAll({ page, perPage }) {
      const body = await api.all({
        page: page,
        perPage: perPage,
        options: { auth: auth.user.token, resp: 'json' },
      })

      if (body.infos) {
        const owners = await users.get({ ids: body.infos.map((m) => m.owner_id) })

        body.infos.forEach((w) => w.owner_name = owners.find((o) => o.id == w.owner_id)?.username)
      }

      runInAction(() => {
        this.models.infos = body.infos
        this.models.total = body.total
      })

      localStorage.setItem(modelsKey, JSON.stringify(this.models));

      return this.models.infos
    },
    async get({ id }) {
      return await api.one({
        resourceID: id,
        options: { auth: auth.user.token, resp: 'json' },
      })
    },
    async download({ ids }) {
      const zip = new JSZip();
      const models = zip.folder('models');

      const m = await Promise.all(ids.map((id) => this.get({ id: id })))

      m.forEach((model) => {
        const f = models?.folder(model.id)

        const fw = f?.folder('weights')
        model.structure.weights?.forEach((w) => {
          fw?.file(w.id + '.json', JSON.stringify(w))
        })

        delete model.structure.weights;

        f?.file('model.json', JSON.stringify(model))
      })

      zip.generateAsync({ type: "blob" })
        .then(function (content) {
          const a = document.createElement("a");
          a.href = URL.createObjectURL(content);
          a.download = 'models.zip';
          a.click();
        })
    },
    async upload({ title, structure, weights }): Promise<void> {
      await api.create({
        payload: {
          'title': title,
          'structure': structure,
          'weights': weights
        },
        options: { auth: auth.user.token }
      })

      return
    },
    async delete({ ids }): Promise<void[]> {
      return Promise.all(
        ids.map((id) => { api.destroy({ resourceID: id, options: { auth: auth.user.token } }) })
      )
    }
  }
}
