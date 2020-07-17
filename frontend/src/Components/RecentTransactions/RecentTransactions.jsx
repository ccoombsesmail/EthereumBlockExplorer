import React, { useState, useEffect } from 'react'
import RecentTransactionsItem from './RecentTransactionsItem/RecentTransactionsItem'
import styles from './RecentTransactions.module.css'
import axios from 'axios';




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
              <RecentTransactionsItem key={i} time={i} transaction={transaction} isVisible={isVisible} />
            )
          })
        }
      </div>

    )

}


export default RecentTransactions