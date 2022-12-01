import React, { FormEvent, useState } from 'react';

import 'bulma/css/bulma.min.css';
import { Button } from 'react-bulma-components';

import { faEnvelope, faUser, faKey, faEye, } from '@fortawesome/free-solid-svg-icons'

import { useAuthStore } from 'contexts/authContext';
import InputField from 'components/Fields/InputField';
import { toast } from 'react-toastify';

type Errors = {
  email?: string,
  username?: string,
  fullname?: string,
  password?: string,
  confirmation?: string
}

function RegistrationForm() {
  const auth = useAuthStore();

  const [state, setState] = useState({
    email: "", username: "", fullname: "", password: "", confirmation: ""
  });
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState(state => ({ ...state, [e.target.name]: e.target.value }))
  }

  const [loading, setLoading] = useState(false);

  const [passwordVisibility, setPasswordVisibility] = useState(false);
  const [confirmationVisibility, setConfirmationVisibility] = useState(false);

  const [errors, setErrors] = useState<Errors>({})

  const validate = () => {
    setErrors({})

    let ok = true
    if (!state.email) {
      setErrors((errors) => ({ ...errors, email: "*required" }))
      ok = false
    }

    if (!state.username) {
      setErrors((errors) => ({ ...errors, username: "*required" }))
      ok = false
    }

    if (!state.fullname) {
      setErrors((errors) => ({ ...errors, fullname: "*required" }))
      ok = false
    }

    if (!state.password) {
      setErrors((errors) => ({ ...errors, password: "*required" }))
      ok = false
    }

    if (!state.confirmation) {
      setErrors((errors) => ({ ...errors, confirmation: "*required" }))
      ok = false
    }

    if (state.password != state.confirmation) {
      setErrors((errors) => ({
        ...errors,
        password: "*doesn't match",
        confirmation: "*doesn't match"
      }))
      ok = false
    }

    return ok
  }

  const register = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    e.stopPropagation()

    if (!validate())
      return

    setLoading(true);
    auth.register({ ...state })
      .catch((err) => toast(err.message))
      .finally(() => setLoading(false));
  }

  return (
    <form
      onSubmit={register}>

      <InputField
        lable={"Email"}
        placeholder={"Email"}
        name={"email"}
        type={"email"}
        onChange={handleChange}
        icon={faEnvelope}
        value={state.email}
        error={errors.email}
      />
      <InputField
        lable={"Username"}
        placeholder={"Username"}
        name={"username"}
        onChange={handleChange}
        icon={faUser}
        value={state.username}
        error={errors.username}
      />
      <InputField
        lable={"Fullname"}
        placeholder={"John Smith"}
        name={"fullname"}
        onChange={handleChange}
        icon={faUser}
        value={state.fullname}
        error={errors.fullname}
      />
      <InputField
        lable={"Password"}
        placeholder={"Password"}
        name={"password"}
        type={passwordVisibility ? "text" : "password"}
        onChange={handleChange}
        icon={faKey}
        rightIcon={faEye}
        onClick={() => { setPasswordVisibility(!passwordVisibility) }}
        value={state.password}
        error={errors.password}
      />
      <InputField
        lable={"Confirm Password"}
        placeholder={"Password"}
        name={"confirmation"}
        type={confirmationVisibility ? "text" : "password"}
        onChange={handleChange}
        icon={faKey}
        rightIcon={faEye}
        onClick={() => { setConfirmationVisibility(!confirmationVisibility) }}
        value={state.confirmation}
        error={errors.confirmation}
      />

      <Button.Group>
        <Button
          mt="3"
          fullwidth
          rounded
          color="primary"
          loading={loading}
        >
          Sign Up
        </Button>
      </Button.Group>
    </form>
  )
}

export default RegistrationForm;
