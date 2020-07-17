import React from 'react'
import styles from './Header.module.css'
import SearchBar from './SearchBar/SearchBar'


const Header = () => {
  return (
    <header className = {styles.headerWrapper}>
      <div className = {styles.left}>
        Block Explorer
      </div>
      <SearchBar/>
      <div className={styles.right}>
        Info
      </div>
    </header>
  )
}



export default Header
