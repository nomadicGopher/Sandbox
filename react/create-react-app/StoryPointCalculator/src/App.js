import React, { useState } from 'react';
import './App.css';

function App() {
	const [complexity, setComplexity] = useState(1);
	const [time, setTime] = useState(1);
	const value = complexity * time;

	return (
		<div className="container-fluid vh-100 position-relative">
			<form className="position-absolute top-50 start-50 translate-middle">
				<div className="row">
					<h2>Complexity</h2>
					<div className="col">
						<input id="complexity-easy" type="radio" value="1" name="complexity" aria-label="Easy level of complexity" checked={complexity === 1} onChange={() => setComplexity(1)} /><br />
						<label htmlFor="complexity-easy" className="form-label">Easy</label>
					</div>
					<div className="col">
						<input id="complexity-moderate" type="radio" value="2" name="complexity" aria-label="Moderate level of complexity" checked={complexity === 2} onChange={() => setComplexity(2)} /><br />
						<label htmlFor="complexity-moderate" className="form-label">Moderate</label>
					</div>
					<div className="col">
						<input id="complexity-difficult" type="radio" value="3" name="complexity" aria-label="Difficult level of complexity" checked={complexity === 3} onChange={() => setComplexity(3)} /><br />
						<label htmlFor="complexity-difficult" className="form-label">Difficult</label>
					</div>
				</div>
				&nbsp;
				<div className="row">
					<h2>Time</h2>
					<div className="col">
						<input id="time-day" type="radio" value="1" name="time" aria-label="1 day or less of effort" checked={time === 1} onChange={() => setTime(1)} /><br />
						<label htmlFor="time-day" className="form-label">1 Day or Less</label>
					</div>
					<div className="col">
						<input id="time-week" type="radio" value="2" name="time" aria-label="2 - 5 days of effort" checked={time === 2} onChange={() => setTime(2)} /><br />
						<label htmlFor="time-week" className="form-label">2-5 Days</label>
					</div>
					<div className="col">
						<input id="time-bi" type="radio" value="4" name="time" aria-label="1 - 2 weeks of effort" checked={time === 4} onChange={() => setTime(4)} /><br />
						<label htmlFor="time-bi" className="form-label">1-2 Weeks</label>
					</div>
					<div className="col">
						<input id="time-month" type="radio" value="6" name="time" aria-label="2 weeks - 1 month of effort" checked={time === 6} onChange={() => setTime(6)} /><br />
						<label htmlFor="time-month" className="form-label">2-4 Weeks</label>
					</div>
				</div>
				<div className="row">
					<div className="col">
						<a id="total"><span id="value">{value}</span> Story Point{value !== 1 ? 's' : ''}</a><br />
						<button type="submit" className="invisible">_</button>
					</div>
				</div>
			</form>
		</div>
	);
}

export default App;
