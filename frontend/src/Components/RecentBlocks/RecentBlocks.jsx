import React from 'react'
import axios from 'axios';
import RecentBlocksItem from './RecentBlocksItem/RecentBlocksItem'
import styles from './RecentBlocks.module.css'

import "animate.css/animate.min.css";
import { connect } from "../util/websocket_util";
import RecentTransactions from '../RecentTransactions/RecentTransactions';



class RecentBlocks extends React.Component {

  constructor() {
    super()
    this.state = {
      blocks: [],
      isVisible: true
    }

    this.loadRecentBlocks = this.loadRecentBlocks.bind(this)
    this.loadOutBlocks = this.loadOutBlocks.bind(this)
    this.blocksQueue = []
  }


  componentDidMount() {
    this.loadRecentBlocks()
    this.socket = connect()
    this.socket.onmessage = msg => {
      this.blocksQueue.push(JSON.parse(msg.data))
      console.log(this.blocksQueue)
      if (this.blocksQueue.length === 4) {
        let newBlocks = [...this.blocksQueue]
        this.blocksQueue = []
        this.setState({ isVisible: false })
        setTimeout(() => {
          this.setState({ blocks: newBlocks, isVisible: true })
        }, 2000)
      }
    };

  }

  componentWillUnmount(){
    this.socket.close()
  }


  async loadRecentBlocks() {
    try {
      const res = await axios.get('https://eth-blockexplorer-go.uc.r.appspot.com/api/recentblocks')
      this.setState({blocks: res.data, isVisible: true})
    } catch (e) {
      console.log(e)
    }
  }

  loadOutBlocks() {
    this.setState({isVisible: false})
  }




  render() {
    const { blocks } = this.state
    return (
      <div className = {styles.wrapper}>
          <h1>Recent Blocks</h1>
        <div className={styles.blocksWrapper}>
          {
            blocks.map((block,i) => {
              return (
                <RecentBlocksItem key={i} time={i} block={block} isVisible={this.state.isVisible} />
              )
            })
          }
        </div>
        <div>
          <h1>Recent Transactions</h1>
          <RecentTransactions/>
        </div>
      </div>
    )
  }

}


export default RecentBlocks