import React, { Dispatch, FormEvent, SetStateAction, useState } from 'react';

import 'bulma/css/bulma.min.css';
import { Button } from 'react-bulma-components';

import { faPen } from '@fortawesome/free-solid-svg-icons'

import { toast } from 'react-toastify';

import { useModelsStore } from 'contexts/modelsContext';

import InputField from 'components/Fields/InputField';
import InputFile from 'components/Fields/InputFile';

function UploadModelForm(
  { setInProgress }: { setInProgress: Dispatch<SetStateAction<boolean>> }
) {
  const models = useModelsStore();

  const [state, setState] = useState<{
    title: string, structure?: File | null, weights?: File | null
  }>({
    title: "", structure: null, weights: null
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState(state => ({ ...state, [e.target.name]: e.target.value }))
  }

  const handleStructureChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState(state => ({ ...state, structure: e.target.files?.item(0) }))
  }

  const handleWeightsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState(state => ({ ...state, weights: e.target.files?.item(0) }))
  }

  const [loading, setLoading] = useState(false);

  const [errors, setErrors] = useState<{
    title?: string, structure?: string, weights?: string
  }>({})

  const validate = () => {
    setErrors({})

    let ok = true
    if (!state.title) {
      setErrors((errors) => ({ ...errors, title: "*required" }))
      ok = false
    }

    if (!state.structure) {
      setErrors((errors) => ({ ...errors, structure: "*required" }))
      ok = false
    }

    if (!state.weights) {
      setErrors((errors) => ({ ...errors, weights: "*required" }))
      ok = false
    }

    return ok
  }

  const upload = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    e.stopPropagation()

    if (!validate())
      return

    if (!state.structure || !state.weights)
      return

    setLoading(true);

    models.upload(
      { title: state.title, structure: state.structure, weights: state.weights }
    )
      .catch((err) => toast(err.message))
      .then(() => setInProgress(false))
      .finally(() => {
        setLoading(false)
      });
  }

  return (
    <form onSubmit={upload}>
      <InputField
        lable={"Title"}
        placeholder={"Title"}
        name={"title"}
        onChange={handleChange}
        icon={faPen}
        value={state.title}
        error={errors.title}
      />

      <InputFile
        lable={"Structure"}
        placeholder={state.structure?.name ?? "Choose structure file..."}
        name={"email"}
        onChange={handleStructureChange}
        error={errors.structure}
      />

      <InputFile
        lable={"Weights"}
        placeholder={state.weights?.name ?? "Choose weights file..."}
        name={"email"}
        onChange={handleWeightsChange}
        error={errors.weights}
      />

      <Button.Group align='right'>
        <Button
          mt="3"
          rounded
          color="success"
          loading={loading}
        >
          Submit
        </Button>
      </Button.Group>
    </form>
  )
}

export default UploadModelForm;
