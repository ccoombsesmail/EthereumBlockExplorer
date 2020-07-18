import React, { useState } from 'react'
import styles from './SearchBar.module.css'
import { withRouter } from 'react-router-dom'
const SearchBar = (props) => {

  const [input, setInput] = useState("")

  const handleSubmit = (e) => {
    e.preventDefault()
    // axios.get('/api/transactions').then((res) => console.log(res))
    // axios.get('/api/transactions', {
    //   params: {
    //     searchQuery: input
    //   }
    // }).then((res) =>{ 
    //   console.log(res)
    // })
    // const location = {
    //   pathname: `/transactions/${}`,
    //   state: { fromDashboard: true }
    // }
    props.history.push(`/transactions/${input}`)
    setInput("")
  }

  const handleUpdate = (e) => {
    setInput(e.currentTarget.value)
  }

  return (
    <form className={styles.searchBarWrapper} onSubmit={handleSubmit} spellCheck="false">
        <input onChange={handleUpdate} type="text" value={input} placeholder="Enter Tx Hash, Sender Address, Receiver Address..."/>
    </form>
  )
}




export default withRouter(SearchBar)