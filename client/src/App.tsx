import { useState } from 'react'
import './App.css'

function App() {
  const [fact, setFact] = useState<string>("")

  const handleClick = async () => {
    try {
        const res = await fetch("http://localhost:3000");
        const data = await res.json();
        setFact(data.fact);
    } catch (error) {
        setFact("Error while getting the cat fact!")
    }
}

  return (
    <main className='container'>
      <h2>Get Cat Fact!</h2>
      <button onClick={handleClick}>See one</button>
      <p>{fact}</p>
    </main>
  )
}

export default App
