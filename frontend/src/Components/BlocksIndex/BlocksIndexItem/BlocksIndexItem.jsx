import React from 'react'
import styles from './BlocksIndexItem.module.css'
import "animate.css/animate.min.css";
import { Animated } from "react-animated-css";

const BlocksIndexItem = ({block, time}) => {
  console.log(block)
  return (
    <Animated animationIn="fadeInUpBig" animationInDuration={400} animationInDelay={time * 40}  >
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
        <div className={styles.blockTransactions}>
          <h3 className={styles.transactionNumber}>
            {/* {block.Transactions.length}  */}
            Transactions
          </h3>
        </div>
      </div>
    </Animated>

  )
  
}

export default BlocksIndexItem