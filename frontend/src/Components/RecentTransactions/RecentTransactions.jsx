import React, { useState, useEffect } from 'react'
import RecentTransactionsItem from './RecentTransactionsItem/RecentTransactionsItem'
import styles from './RecentTransactions.module.css'
import axios from 'axios';
import { Link } from 'react-router-dom'



const RecentTransactions = () => {

  const [isVisible, setIsVisible] = useState(true)
  const [transactions, setTransactions] = useState([])
  // setInterval(loadRecentTransactions(), 10000)


  function wait(time) {
    return new Promise(resolve => {
      setTimeout(() => {
        resolve('');
      }, time);
    });
  }

  const loadRecentTransactions = async() => {
    try {
      const res = await axios.get('/api/recentblocks')
      setIsVisible(true)
      setTransactions(res.data[0].Transactions.slice(0, 20))
      await wait(10000).then(() => setIsVisible(false))
      await wait(2000)
      return await loadRecentTransactions() 
    } catch (e) {
      console.log(e)
      loadRecentTransactions() 
    }
  }



  useEffect(() => {
    loadRecentTransactions()
  }, [])

    return (
      <div className={styles.transactionsWrapper}>
        {
          transactions.map((transaction, i) => {
            return (
              <Link key={i} to={{
                pathname: `/transaction/${transaction.Hash}`,
                state: { transaction: transaction }
              }}> 
                <RecentTransactionsItem  time={i} transaction={transaction} isVisible={isVisible} />
              </Link>
            )
          })
        }
      </div>

    )

}


export default RecentTransactions