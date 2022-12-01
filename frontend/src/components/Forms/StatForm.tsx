import { FormEvent, useState } from 'react';

import 'bulma/css/bulma.min.css';
import { Button } from 'react-bulma-components';

import { useStatsStore } from "contexts/statsContext";

import InputField from 'components/Fields/InputField';

import { toast } from 'react-toastify';

type StatFormParams = {
  object: 'users' | 'weights' | 'models'
}

function StatForm({ object }: StatFormParams) {
  const stats = useStatsStore();

  const [state, setState] = useState({ from: "", to: "" });
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState(state => ({ ...state, [e.target.name]: e.target.value }))
  }

  const [loading, setLoading] = useState(false);

  const [errors, setErrors] = useState<
    { from?: string, to?: string }>({ from: "", to: "" });

  const validate = () => {
    setErrors({})

    let ok = true
    if (!state.from) {
      setErrors((errors) => ({ ...errors, from: "*required" }))
      ok = false
    }

    if (!state.to) {
      setErrors((errors) => ({ ...errors, to: "*required" }))
      ok = false
    }

    if (new Date(state.from) > new Date(state.to)) {
      setErrors((errors) => ({ ...errors, to: "*invalid" }))
      ok = false
    }

    return ok
  }

  const submit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    e.stopPropagation()

    if (!validate()) {
      return
    }

    setLoading(true);
    (() => {
      switch (object) {
        case 'users':
          return stats.getUsers({ from: new Date(state.from), to: new Date(state.to) });
        case 'weights':
          return stats.getWeights({ from: new Date(state.from), to: new Date(state.to) });
        case 'models':
          return stats.getModels({ from: new Date(state.from), to: new Date(state.to) });
        default:
          throw new Error(`Error: unsupported object: ${object}`)
      }
    })()
      .catch((err) => toast(err.message))
      .finally(() => setLoading(false));
  }

  return (
    <form onSubmit={submit} >
      <InputField
        lable={"From"}
        name={"from"}
        type={"datetime-local"}
        onChange={handleChange}
        value={state.from}
        error={errors.from}
      />
      <InputField
        lable={"To"}
        name={"to"}
        type={"datetime-local"}
        onChange={handleChange}
        value={state.to}
        error={errors.to}
      />

      <Button.Group>
        <Button
          fullwidth
          rounded
          color="primary"
          loading={loading}
        >
          Submit
        </Button>
      </Button.Group>
    </form >
  )
}

export default StatForm;
