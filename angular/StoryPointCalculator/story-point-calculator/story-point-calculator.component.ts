import { Component } from '@angular/core';

@Component({
	selector: 'app-story-point-calculator',
	templateUrl: './story-point-calculator.component.html',
	styleUrls: ['./story-point-calculator.component.css']
})
export class StoryPointCalculatorComponent {
	complexity: number = 1;
	time: number = 1;
	value: number = 1;

	calculate() {
		this.value = this.time * this.complexity;
	}
}
