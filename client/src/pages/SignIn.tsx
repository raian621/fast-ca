import { createRef, FormEvent } from 'react'
import { z, ZodError } from 'zod'

const formSchema = z.object({
  username: z.string().min(5).max(256),
  email: z.string().email().min(5).max(256),
  password: z.string().min(10),
})

export default function SignIn() {
  const handleSubmit = (e: FormEvent) => {
    e.preventDefault()

    const username = fieldRefs['username'].current?.value
    const email = fieldRefs['email'].current?.value
    const password = fieldRefs['password'].current?.value
    const passwordConfirm = fieldRefs['passwordConfirm'].current?.value

    if (password != passwordConfirm) {
      return
      //setErrorMsg("Passwords must match");
    }
    const formData = { username, email, password }
    try {
      formSchema.parse(formData)
    } catch (err) {
      if (err instanceof ZodError) {
        console.error(err.issues)
      }
    }

    console.log(formData)
  }

  const fieldRefs = {
    username: createRef<HTMLInputElement>(), 
    email: createRef<HTMLInputElement>(), 
    password: createRef<HTMLInputElement>(), 
    passwordConfirm: createRef<HTMLInputElement>(), 
  }

  return (
    <>
      <h1>
        Sign in
      </h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor='username'>Username</label>
        <input ref={fieldRefs['username']} type='text' name='username'/>
        <label htmlFor='email'>Email</label>
        <input ref={fieldRefs['email']} type='email' autoComplete='email' name='email'/>
        <label htmlFor='password'>Password</label>
        <input ref={fieldRefs['password']} type='password' name='password'/>
        <label htmlFor='passwordConfirm'>Confirm password</label>
        <input ref={fieldRefs['passwordConfirm']} type='password' name='passwordConfirm'/>
        <input type='submit'/>
      </form>
    </>
  )
}
