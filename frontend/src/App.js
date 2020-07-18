import React from 'react';
import './App.css';
import BlockIndex from './Components/BlocksIndex/BlocksIndex'
import RecentBlocks from './Components/RecentBlocks/RecentBlocks'
import SideBar from './Components/SideBar/SideBar';
import Header from './Components/Header/Header'
import { Route, Switch } from 'react-router-dom'
import TransactionsIndex from './Components/TransactionsIndex/TransactionsIndex';
import TransactionShow from './Components/TransactionShow/TransactionShow';

function App() {
  return (
    <div className="App">
      <Header/>
      <div className = "mainContent">
        <SideBar/>
        <Switch>
          <Route path="/index" component={BlockIndex} />  
          <Route path="/transactions/:search" component={TransactionsIndex} />  
          <Route path="/transaction/:hash" component={TransactionShow} />  
          <Route path="/" component={RecentBlocks} />  
        </Switch>
      </div>
    </div>
  );
}

export default App;
