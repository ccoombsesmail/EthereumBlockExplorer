import React from 'react'
import axios from 'axios';
import BlocksIndexItem from './BlocksIndexItem/BlocksIndexItem'
import styles from './BlocksIndex.module.css'

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
            return <BlocksIndexItem time = {i} block = {block}/>
          })
        }
      </div>
    )
  }

}


export default BlocksIndex