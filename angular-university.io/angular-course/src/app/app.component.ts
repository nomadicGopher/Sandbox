import { Component } from '@angular/core';
//? import { RouterOutlet } from '@angular/router';

@Component({  //! Component aka Controller
  selector: 'app-root',
  //? imports: [RouterOutlet],
  templateUrl: './app.component.html', //! Template aka View
  styleUrl: './app.component.css'
})
export class AppComponent { //! Data aka Model
  data = {
    title: 'angular-course'
  };

  onLogoClicked() {
    alert('Hello World');
  }

  onKeyUp(newTitle:string) {
    //! onKeyUp() takes newTitle as an input value dynamically being updated by the (keyup) method. The goal is to reassign the new string value to data.Title to dynamically update the title.
    this.data.title = newTitle;
  }
}
