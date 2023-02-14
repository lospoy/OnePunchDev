import { Box, List, ThemeIcon } from '@mantine/core'
import { CheckCircleFillIcon } from '@primer/octicons-react'
import useSWR from 'swr'
import './App.css'
import AddTodo from './components/AddTodo'

export interface Todo {
  id: number
  title: string
  body: string
  done: boolean
}

// Defining endpoint
export const ENDPOINT = 'http://localhost:4000'

// Creating our own fetcher (needed with SWR)
const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then(res => res.json())

function App() {
  // mutate is returned from useSWR
  const {data, mutate} = useSWR<Todo[]>('api/todos', fetcher)

  async function markTodoDone(id: number) {
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/done`, {
      method: 'PATCH',
    }).then(res => res.json())

    mutate(updated)
  }

  return (
    <Box
      sx={(theme) => ({
       padding: '2rem',
       width: '100%',
       maxWidth: '40rem',
       margin: '0 auto',
      })}
    >
      <List spacing='xs' size='sm' mb={12}>
        {data?.map((todo) => {
          // adding "todo_list__" prefix to ensure the key is unique
          return (
            <List.Item
              onClick={() => markTodoDone(todo.id)}
              key={`todo_list__${todo.id}`}
              icon={
                todo.done
                ? (<ThemeIcon color='teal' size={24} radius='xl'><CheckCircleFillIcon size={20}/></ThemeIcon>)
                : (<ThemeIcon color='gray' size={24} radius='xl'><CheckCircleFillIcon size={20}/></ThemeIcon>)
              }
            
            >{todo.title}</List.Item>
          )
        })}
      </List>
      <AddTodo mutate={mutate} />
    </Box>
  )
}

export default App
