import React from 'react'
import axios from 'axios';
import BlocksIndexItem from './BlocksIndexItem/BlocksIndexItem'
import styles from './BlocksIndex.module.css'
import { Link } from 'react-router-dom'
class BlocksIndex extends React.Component {

  constructor() {
    super()
    this.state = {
      blockData: []
    }
  }


  componentDidMount() {
    axios.get('/api/blocks').then(block => {
      this.setState({blockData: block.data})
    }) 
  }


  render() {
    const {blockData} = this.state
    return (
      <div className = {styles.blocksWrapper}>
        {
          blockData.map((block, i) => {
            return ( 
              <Link key={i} to={{
                pathname: `/block/${block.Hash}`,
                state: { block: block }
              }}>
                <BlocksIndexItem key={i} time = {i} block = {block}/>
              </Link>
            )
          })
        }
      </div>
    )
  }

}


export default BlocksIndex