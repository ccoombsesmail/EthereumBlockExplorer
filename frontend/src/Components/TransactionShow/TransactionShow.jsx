
import React, { useState, useEffect } from 'react'
import styles from './TransactionShow.module.css'

import axios from 'axios'
import "animate.css/animate.min.css";
import { Animated } from "react-animated-css";
import { useLocation } from 'react-router-dom'
import { FaMicroscope } from 'react-icons/fa';


const TransactionShow = ({match}) => {
  const [transaction, setTransaction] = useState(null)
  const [gasUsed, setgasUsed] = useState(null)
  let location = useLocation();
  useEffect(() => {
    getGatUsed(match.params.hash)
    if (location.state) {
      let tx = location.state.transaction
      setTransaction(tx)
    } else {
      axios.get('/api/transactions', {
        params: {
          searchQuery: match.params.hash
        }
      }).then((res) => {
        setTransaction(res.data[0])
      })    
    }
  }, [])

  const getGatUsed = (hash) => {
    axios.get('/api/gasused', {
      params: {
        txhash: hash
        }
    }).then((res) => {
      setgasUsed(res.data)
    })
  }


  return (
    <div className={styles.transactionWrapper}>
      
      <Animated animationIn="fadeInUp" animationOut="fadeOutDown" animationInDuration={1500} animationInDelay={300} className={styles.animationWrapper} >
        <FaMicroscope style={{ fontSize: '50px', color: '#21FA90' }} />
        <h1>Transactions Details</h1>
          { 
          transaction !== null ? (
          <ul className={styles.transactionInfo}>
            <li>
              <span>TxHash:</span>{transaction.Hash}
            </li>
            <li>
                <span>From:</span>{transaction.From}
            </li> 
            <li>
                <span>To:</span>{transaction.To}
            </li>
            <li>
              <span>Value:</span> {(Number(transaction.Value) / 1000000000000000000).toString() + " ETH"}
            </li>
            <li>
              <span>Gas Limit:</span>{transaction.GasLimit}
            </li>  
            <li>
              <span>Gas Used:</span>{gasUsed}
            </li>
            <li>
              <span>Size:</span>{transaction.Size}
            </li>
          </ul> ) : null
          }
      </Animated>
    </div>
  )
}

export default TransactionShow