import './App.css'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import SignIn from './pages/SignIn'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/'>
          <Route index element={<h1>Home</h1>}/>
          <Route path="signin" element={<SignIn/>}/>
          <Route path="ca" element={<h1>Certificate authorities</h1>}/>
          <Route path="certificates" element={<h1>Certificates</h1>}/>
          <Route path="api-keys" element={<h1>API Keys</h1>}/>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
