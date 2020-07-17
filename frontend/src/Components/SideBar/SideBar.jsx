import React from 'react'
import styles from './SideBar.module.css'
import {Link} from 'react-router-dom'


const SideBar = () => {

  // const [count, setCount] = useState(0);


  return (
    <div className = {styles.sideBarWrapper}>
        <ul className = {styles.sideBarList}>
          <Link to="/index"><li>Last 100 Blocks</li> </Link>
        </ul>
    </div>
  )
}


export default SideBar