## [angular.dev](https://angular.dev/overview)
* [**Playground**](https://angular.dev/playground)
* [**Tutorials**](https://angular.dev/tutorials)
  * [Learn Angular in the browser](https://angular.dev/tutorials/learn-angular)
  * [Build your first app locally](https://angular.dev/tutorials/first-app)
  * [Deferrable views](https://angular.dev/tutorials/deferrable-views)
* **References**
  * [API](https://angular.dev/api)
  * [CLI](https://angular.dev/cli)
  * [Errors](https://angular.dev/errors)
  * [Extended Diagnosis](https://angular.dev/extended-diagnostics)
  * [Configurations](https://angular.dev/reference/configs/file-structure)

## Security Features
* Angular uses type expectations to prevent malisciouse injections by applying proper escape characters when values are being dynamically referenced. IE in angular-course/src/app/app.component.ts ... if data.title were to be overwitten with some other value containing html elements, they would be preceded by escape characters in order to be rendered as a string rather than being executed (ie. adding script tags). 

## Vocabulary

### Component
The basic building block of an Angular application. It consists of a TypeScript class, an HTML template, and optional CSS styles. Components control a part of the UI and can be nested within other components.

```typescript
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'my-angular-app';
}
```

### Decorator
A special type of declaration that can be attached to a class, method, accessor, property, or parameter. Decorators are used to add metadata to classes and other structures. Common decorators include @Component, @NgModule, and @Injectable.

```typescript
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
```

### Module
A container for a cohesive block of code dedicated to an application domain, a workflow, or a closely related set of capabilities. An Angular application is defined by a set of NgModules. The root module is typically called AppModule.

```typescript
@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
```

### Service
A class that provides a specific functionality to be shared across components. Services are typically used for data fetching, business logic, and other operations that need to be reused. Services are often injected into components using dependency injection.

```typescript
@Injectable({
  providedIn: 'root',
})
export class DataService {
  // Service logic
}
```

### Directive
A class that adds behavior to elements in your Angular applications. There are three kinds of directives: components, structural directives (e.g., *ngIf, *ngFor), and attribute directives (e.g., ngClass, ngStyle).

```typescript
<div *ngIf="isVisible">This is conditionally visible</div>
```

### Pipe
A way to transform data in your templates. Pipes are used to format data, such as dates, numbers, and strings. Angular provides several built-in pipes, and you can also create custom pipes.

```html
<p>{{ today | date:'fullDate' }}</p>
```

### Routing
The mechanism for navigating between different views or components in an Angular application. The Angular Router enables navigation from one view to the next as users perform application tasks.

```typescript
const routes: Routes = [
  { path: 'home', component: HomeComponent },
  { path: 'about', component: AboutComponent }
];
```

### Template
The HTML part of a component that defines the view. Templates can include Angular-specific syntax, such as interpolation, directives, and bindings.

```html
<h1>{{ title }}</h1>
```

### Binding
The mechanism for coordinating parts of a template with parts of a component. There are different types of bindings, including property binding, event binding, and two-way binding.

```html
<input [(ngModel)]="username">
```

### Classes
Basic JS & TS feature used to create objects and define their behavior through properties and methods.

In Angular, classes are used extensively to define components, services, and other constructs. Here's a quick refresher on how classes work in TypeScript:

```typescript
class Person {
  name: string;
  age: number;

  constructor(name: string, age: number) {
    this.name = name;
    this.age = age;
  }

  greet() {
    console.log(`Hello, my name is ${this.name} and I am ${this.age} years old.`);
  }
}

const person = new Person('Alice', 30);
person.greet(); // Output: Hello, my name is Alice and I am 30 years old.
```

In Angular, you use classes to define components, services, and other parts of your application. For example, a component class might look like this:

```typescript
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'my-angular-app';
}
```