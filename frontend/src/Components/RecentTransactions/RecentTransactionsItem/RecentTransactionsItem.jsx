
import React from 'react'
import styles from '../RecentTransactions.module.css'

import "animate.css/animate.min.css";
import { Animated } from "react-animated-css";

const RecentTransactionsItem = ({ transaction, time, isVisible }) => {

  return (
    <Animated animationIn="fadeInUp" animationOut="fadeOutDown" animationInDelay={time*150} animationOutDelay={time * 150} isVisible={isVisible} >
      <div className={styles.transaction}>
        <ul>
          <li>
            {"From: " + transaction.From}
          </li>
          <li>
            {"To: " + transaction.To}
          </li>
          <li>
            {"Value: " + (Number(transaction.Value) / 1000000000000000000).toString() + " ETH"}
          </li>
        </ul>
      </div>
    </Animated>

  )

}

export default RecentTransactionsItem