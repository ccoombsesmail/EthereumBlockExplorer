
import React, { useState, useEffect } from 'react'
import styles from './TransactionShow.module.css'

import "animate.css/animate.min.css";
import { Animated } from "react-animated-css";
import { useLocation } from 'react-router-dom'

const TransactionShow = () => {
  const [transaction, setTransaction] = useState(null)
  let location = useLocation();
  console.log(location)
  useEffect(() => {
    setTransaction(location.state.transaction)
  }, [])

  return (
    <Animated animationIn="fadeInUp" animationOut="fadeOutDown" >
      <div className={styles.transactionWrapper}>
        { 
        transaction !== null ? (
        <ul className={styles.transactionInfo}>
          <li>
            {"From: " + transaction.From}
          </li>
          <li>
            {"To: " + transaction.To}
          </li>
          <li>
            {"Value: " + (Number(transaction.Value) / 1000000000000000000).toString() + " ETH"}
          </li>
        </ul> ) : null
        }
      </div>
    </Animated>
  )
}

export default TransactionShow