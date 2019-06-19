import React from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios'

class App extends React.Component {
    state = {X: 1, Y: 1, Op: '+', calculations: []};
    handleSubmit = (event) => {
      event.preventDefault();
      var {Op, X, Y} = event.target
      axios.get(`http://localhost:8081/calculate/${Op.value}/${X.value}/${Y.value}`).then(() => {
        this.updateCalculations();
      });
    }
    componentDidMount(){
      this.updateCalculations();
      var intervalId = setInterval(this.updateCalculations, 3000);
    }

    updateCalculations = () => {
      axios.get(`http://localhost:8081/calculations`)
      .then(res => {
        const calculations = res.data;
        if (calculations)
          this.setState({ calculations: calculations });
      })
    }

    render() {
      return (
        <div className="App">
        <header className="App-header">
        <p>
          Please input an X, Y, and Op Value and Submit
        </p>
        <form onSubmit={this.handleSubmit}>
          <label>
            X:
            <input type="number" name="X" />
          </label>
          <label>
            Y:
            <input type="number" name="Y" />
          </label>
          <label>
            Op:
            <select name="Op" >
              <option value="+">+</option>
              <option value="-">-</option>
              <option value="*">*</option>
              <option value="d">/</option>
            </select>
          </label>
          <input type="submit" value="Submit" />
        </form>
        <ol>
          {
            this.state.calculations.map((calc, i) => {
              return (
                <li key={"calc"+i}>{calc.X} {(calc.Op === "d") ? "/" : calc.Op} {calc.Y} = {calc.Calc}</li>
              )
            }
          )}
        </ol>
      </header>
    </div>
  );
  }
}

export default App;
