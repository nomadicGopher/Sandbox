import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { StoryPointCalculatorComponent } from './story-point-calculator/story-point-calculator.component';

@NgModule({
	declarations: [
		AppComponent,
		StoryPointCalculatorComponent
	],
	imports: [
		BrowserModule,
		FormsModule
	],
	providers: [],
	bootstrap: [AppComponent]
})
export class AppModule { }
