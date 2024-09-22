import { Input } from '@/shared/ui/input';
import { Button } from '@/shared/ui/button';
import { createStore } from 'solid-js/store';
import { RegisterDto, registerDtoSchema } from '@/shared/lib';
import { register } from '@/shared/lib/server';
import { createSignal, Show } from 'solid-js';

export const RegisterForm = () => {
  const [form, setForm] = createStore<RegisterDto>({
    username: '',
    password: '',
    email: '',
    full_name: '',
  })
  const [error, setError] = createSignal("")

  const onSubmit = async () => {
    const {data, error, success} = await registerDtoSchema.safeParseAsync(form)
    console.log(data, error, success)
    if (success) {
      await register(data)
      setError("Успешно! Теперь можно войти на соответствующей вкладке")
    } else {
      const err = Object.values(error.format()).map(value => {
        if (!Array.isArray(value)) {
          return value._errors.join("\n")
        }
      }).join("\n")
      setError(err)
    }
  }

  return (
    <>
      <Input placeholder='Полное имя' value={form.full_name} onChange={e => setForm('full_name', e.target.value)} />
      <Input placeholder='Имя пользователя' value={form.username} onChange={e => setForm('username', e.target.value)} />
      <Input placeholder='Электронная почта' type='email' value={form.email} onChange={e => setForm('email', e.target.value)} />
      <Input placeholder='Пароль' type='password' value={form.password} onChange={e => setForm('password', e.target.value)} />
      <Button class='bg-white' onClick={onSubmit}>Зарегистрироваться</Button>

      <Show when={error()}>
        <pre class='text-xs'>{error()}</pre>
      </Show>
    </>
  )
}