import JSZip from "jszip";

import { ApiCallMethods } from "hooks/apiCall";

import { AuthStore } from "./authStore";

interface StatEntry {
  id: string
}
export interface Stat {
  load: StatEntry[]
  edit: StatEntry[]
  registration: StatEntry[]
}

export interface StatsStore {
  getUsers(params: { from: Date, to: Date }): Promise<Stat>,
  getModels(params: { from: Date, to: Date }): Promise<Stat>,
  getWeights(params: { from: Date, to: Date }): Promise<Stat>,
  download(params: { object: string, stat: Stat }): Promise<void>
}

export function createStatsStore(
  auth: AuthStore, api: ApiCallMethods<Stat>
): StatsStore {
  return {
    async getUsers({ from, to }) {
      const body = await api.stat({
        query: {
          'from': from.toISOString(), 'to': to.toISOString(),
          'registration': '1', 'update': '1'
        },
        options: { auth: auth.user.token, resp: 'json', suffix: 'users/stats' },
      })

      await this.download({ object: 'users', stat: body })

      return body
    },
    async getModels({ from, to }) {
      const body = await api.stat({
        query: {
          'from': from.toISOString(), 'to': to.toISOString(),
          'edit': '1', 'update': '1'
        },
        options: { auth: auth.user.token, resp: 'json', suffix: 'models/stats' },
      })

      await this.download({ object: 'models', stat: body })

      return body
    },
    async getWeights({ from, to }) {
      const body = await api.stat({
        query: {
          'from': from.toISOString(), 'to': to.toISOString(),
          'update': '1', 'load': '1'
        },
        options: { auth: auth.user.token, resp: 'json', suffix: 'weights/stats' },
      })

      await this.download({ object: 'weights', stat: body })

      return body
    },
    async download({ object, stat }) {
      const zip = new JSZip();
      const stats = zip.folder('stats');

      Object.entries(stat).forEach((field) => {
        stats?.file(field[0] + '.json', JSON.stringify(field[1]))
      })

      zip.generateAsync({ type: "blob" })
      .then(function (content) {
        const a = document.createElement("a");
        a.href = URL.createObjectURL(content);
        a.download = object + '_stats.zip';
          a.click();
        })
    }
  }
}
