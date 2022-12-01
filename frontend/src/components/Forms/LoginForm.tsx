import { FormEvent, useState } from 'react';

import 'bulma/css/bulma.min.css';
import { Button } from 'react-bulma-components';

import { faEnvelope, faEye, faKey } from '@fortawesome/free-solid-svg-icons'

import { useAuthStore } from 'contexts/authContext';
import InputField from 'components/Fields/InputField';
import { toast } from 'react-toastify';

function LoginForm() {
  const auth = useAuthStore();

  const [state, setState] = useState({ email: "", password: "" });
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState(state => ({ ...state, [e.target.name]: e.target.value }))
  }

  const [loading, setLoading] = useState(false);

  const [passwordVisibility, setPasswordVisibility] = useState(false);

  const [errors, setErrors] = useState<
    { email?: string, password?: string }>({
      email: "",
      password: ""
    });

  const validate = () => {
    setErrors({})

    let ok = true
    if (!state.email) {
      setErrors((errors) => ({ ...errors, email: "*required" }))
      ok = false
    }


    if (!state.password) {
      setErrors((errors) => ({ ...errors, password: "*required" }))
      ok = false
    }

    return ok
  }

  const authorize = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    e.stopPropagation()

    if (!validate()) {
      return
    }

    setLoading(true);
    auth.authorize(state.email, state.password)
      .catch((err) => toast(err.message))
      .finally(() => setLoading(false));
  }

  return (
    <form onSubmit={authorize} >
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

      <Button.Group>
        <Button
          fullwidth
          rounded
          color="primary"
          loading={loading}
        >
          Sign In
        </Button>
      </Button.Group>
    </form >
  )
}

export default LoginForm;
