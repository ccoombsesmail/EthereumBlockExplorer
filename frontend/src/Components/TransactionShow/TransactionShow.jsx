
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
    if (location.state) {
    setTransaction(location.state.transaction)
    }
  }, [])

  return (
    <div className={styles.transactionWrapper}>
      <Animated animationIn="fadeInUp" animationOut="fadeOutDown" className={styles.animationWrapper} >
          { 
          transaction !== null ? (
          <ul className={styles.transactionInfo}>
            <li>
              {"TxHash: " + transaction.From}
            </li>
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
      </Animated>
    </div>
  )
}

export default TransactionShow