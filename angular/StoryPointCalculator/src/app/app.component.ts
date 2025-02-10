import { Component } from '@angular/core';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.css']
})
export class AppComponent {
	complexity: number = 1;
	time: number = 1;
	value: number = 1;
	plural: string = '';

	calculate() {
		this.value = this.time * this.complexity;
		this.plural = this.value !== 1 ? 's' : '';
	}
}
