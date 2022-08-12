import { useForm, SubmitHandler } from "react-hook-form";

type Inputs = {
  userID: string;
  password: string;
}

function App() {

  const { register, handleSubmit, formState: { errors } } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = data => console.log(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input {...register('userID')} placeholder='user id' />
      {errors?.userID && <p>{errors.userID.message}</p>}
      <input {...register('password')} placeholder='password' />
      <input type='submit' />
    </form>
  )
}

export default App
