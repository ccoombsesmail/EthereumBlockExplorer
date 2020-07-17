import React, { useState, useEffect } from 'react'
import TransactionsItem from '../RecentTransactions/RecentTransactionsItem/RecentTransactionsItem'
import styles from './TransactionsIndex.module.css'
import axios from 'axios';
import {withRouter} from 'react-router-dom'




const TransactionsIndex = (props) => {

  const [transactions, setTransactions] = useState([])

  useEffect(() => {
    loadRecentTransactions()
  }, [props.match.params.search])

  const loadRecentTransactions = () => {
    axios.get('/api/transactions', {
      params: {
        // searchQuery: props.location.state['searchQuery']
        searchQuery: props.match.params.search
      }
    }).then((res) => {
      console.log(res)
      setTransactions(res.data)
    })
  }


  return (
    <div className={styles.transactionsWrapper}>
      {
        transactions.map((transaction, i) => {
          return (
            <TransactionsItem key={i} time={i} transaction={transaction} isVisible={true} />
          )
        })
      }
    </div>

  )

}


export default withRouter(TransactionsIndex)