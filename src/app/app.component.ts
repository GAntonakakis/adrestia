import { Component, OnInit } from '@angular/core';
import { ApiService } from './api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
  message: string ='';
  firstname: string = '';
  lastname: string = '';
  age: string = '';
  gender: string = '';
  ethnicity: string = '';
  listNotEmpty: boolean = false;
  listElement: Array<any> = [];

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.apiService.getMessage().subscribe((response) => {
      this.message = response.message;
    });
  }

  sendData(): void {
    const data = {firstname: this.firstname, lastname: this.lastname, age: this.age, gender: this.gender, ethnicity: this.ethnicity};
    this.apiService.sendData(data).subscribe((response) => {
      this.message = response.message;
    });

    //Make the Users list be hidden if empty. Push all new entries in the list.
    this.listNotEmpty = true;
    this.listElement.push({
      'name': this.firstname + " " + this.lastname,
      'age': this.age,
      'gender': this.gender,
      'ethnicity': this.ethnicity
    });

    //Resetting of the Input Text fields after Submitting
    this.firstname = '';
    this.lastname = '';
    this.age = '';
    this.gender = '';
    this.ethnicity = '';
  }
}