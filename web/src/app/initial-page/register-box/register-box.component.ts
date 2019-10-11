import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-register-box',
  templateUrl: './register-box.component.html',
  styleUrls: ['./register-box.component.scss']
})
export class RegisterBoxComponent implements OnInit {
  nameGroup: FormGroup;
  emailGroup: FormGroup;
  usernameGroup: FormGroup;
  passwordGroup: FormGroup;

  constructor(private _formBuilder: FormBuilder) {}

  ngOnInit() {
    this.nameGroup = this._formBuilder.group({
      name: ['', Validators.required]
    });
    this.emailGroup = this._formBuilder.group({
      email: ['', Validators.required]
    });
    this.usernameGroup = this._formBuilder.group({
      username: ['', Validators.required]
    });
    this.passwordGroup = this._formBuilder.group({
      password: ['', Validators.required]
    });
  }

  register() {
    
  }
}
