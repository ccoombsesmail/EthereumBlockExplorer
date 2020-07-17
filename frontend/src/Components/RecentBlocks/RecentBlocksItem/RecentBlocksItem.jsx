import React from 'react'
import styles from './RecentBlocksItem.module.css'

import "animate.css/animate.min.css";
import { Animated } from "react-animated-css";

const RecentBlocksItem = ({ block, time, isVisible }) => {

  let transactionsLen = 0
  if (block.Transactions) {
    transactionsLen = block.Transactions.length
  }
  return (
    <Animated animationIn="bounceInRight" animationOut="flipOutX" animationInDuration={800} animationInDelay={time * 150} animationOutDelay={time*150} isVisible ={isVisible} >
      <div className={styles.block}>
        <ul>
          <li>
            <span>Hash:</span>{" " + block.Hash.slice(0, 8) + '...' + block.Hash.slice(60)}
          </li>
          <li>
            <span>Block Number:</span> {" " + block.Number}
          </li>
          <li>
            <span>Nonce:</span> {" " + block.Nonce}
          </li>
        </ul>
        <div className = {styles.blockTransactions}>
          <h3 className = {styles.transactionNumber}>
            {transactionsLen} Transactions
          </h3>
        </div>
      </div>
    </Animated>

  )

}

export default RecentBlocksItem