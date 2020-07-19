
import React, { useState, useEffect } from 'react'
import styles from './BlockShow.module.css'

import axios from 'axios'
import "animate.css/animate.min.css";
import TransactionsItem from '../RecentTransactions/RecentTransactionsItem/RecentTransactionsItem'
import { useLocation, Link } from 'react-router-dom'


const BlockShow = ({ match }) => {
  const [block, setBlock] = useState(null)
  let location = useLocation();
  useEffect(() => {
    if (location.state) {
      let block = location.state.block
      setBlock(block)
    } else {
      axios.get('/api/block', {
        params: {
          hash: match.params.hash
        }
      }).then((res) => {
        setBlock(res.data)
      })
    }
  }, [])

  return (
    <div className={styles.blockShowWrapper}>

        <h1>Block Details</h1>
        {
          block !== null ? (
            <>
            <ul className={styles.blockInfo}>
              <li>
                <span>Hash:</span>{block.Hash}
              </li>
              <li>
                <span>Block Number:</span>{block.Number}
              </li>
              <li>
                <span>Nonce:</span>{block.Nonce}
              </li>
            </ul> 
            <hr />
          <div className={styles.transactionsWrapper}>
            <h1>Block Transactions</h1>
            {
              block.Transactions.map((transaction, i) => {
                return (
                  <Link key={i} to={{
                    pathname: `/transaction/${transaction.Hash}`,
                    state: { transaction: transaction }
                  }}>
                    <TransactionsItem time={i/2} transaction={transaction} isVisible={true} />
                  </Link>
                )
              })
            }
          </div>
          </>
        ) : null
      }
    </div>
  )
}

export default BlockShow