import React from 'react'
import styles from './Header.module.css'
import SearchBar from './SearchBar/SearchBar'
import { Link } from 'react-router-dom'
// import { FaCubes } from 'react-icons/fa';
import { FaSearch } from 'react-icons/fa';

const Header = () => {
  return (
    <header className = {styles.headerWrapper}>
      <div className = {styles.left}>
        <Link to="/"> ETH Block Scout <FaSearch/> </Link>
      </div>
      <SearchBar/>
      <div className={styles.right}>
        <ul className={styles.linkWrapper}>
          <li>
            <a href="https://www.linkedin.com/in/charles-coombs-esmail-1b5a20113/">
              <img className={styles.icon} src="https://i.ibb.co/28km227/linkedin-icon.png" alt="" />
            </a>
          </li>
          <li>
            <a href="https://github.com/ccoombsesmail">
              <img className={styles.icon} src="https://i.ibb.co/r7dzbWg/github-icon.png" alt="" />
            </a>
          </li>
          <li>
            <a href="https://angel.co/u/charlie-coombs">
              <img className={styles.icon} src="https://i.ibb.co/k1qWrsg/angelist.png" alt="" />
            </a>
          </li>
        </ul>
      </div>
    </header>
  )
}



export default Header
