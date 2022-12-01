import React, { Dispatch, FormEvent, SetStateAction, useState } from 'react';

import { Observer } from 'mobx-react';

import 'bulma/css/bulma.min.css';
import { Button } from 'react-bulma-components';

import { faPen } from '@fortawesome/free-solid-svg-icons';

import { toast } from 'react-toastify';

import { useWeightsStore } from 'contexts/weightsContext';

import InputFile from 'components/Fields/InputFile';
import InputSelect from 'components/Fields/InputSelect';
import InputField from 'components/Fields/InputField';

interface EditModelFormProps {
  modelID: string
  setInProgress: Dispatch<SetStateAction<boolean>>
}

function EditModelForm(
  { modelID, setInProgress }: EditModelFormProps
) {
  const weights = useWeightsStore()

  const [state, setState] = useState<{
    title: string, weights?: File | null
  }>({
    title: "", weights: null
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState(state => ({ ...state, [e.target.name]: e.target.value }))
  }

  const [selectState, setSelectState] = useState<string>("")

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

    if (!state.weights)
      return;

    setLoading(true);
    (!selectState ?
      weights.add({
        modelID: modelID,
        title: state.title,
        weights: state.weights
      }) :
      weights
        .update({
          modelID: modelID,
          id: selectState,
          title: state.title || selectState,
          weights: state.weights
        }))
      .catch((err) => toast(err.message))
      .then(() => setInProgress(false))
      .finally(() => { setLoading(() => false) });
  }

  return (
    <form onSubmit={upload}>
      <InputSelect
        lable="Weight to edit"
        placeholder="Add a new one"
        onChange={(e) => { setSelectState(() => e.target.value) }}
      >
        <Observer>
          {() =>
            <>
              {weights.weights?.map((w) =>
                <option key={w.id} value={w.id}>{w.name}</option>
              )}
            </>
          }
        </Observer>
      </InputSelect>

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

export default EditModelForm;
