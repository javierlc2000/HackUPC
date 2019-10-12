import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginBoxComponent } from './initial-page/login-box/login-box.component';
import { RegisterBoxComponent } from './initial-page/register-box/register-box.component';
import { CentralContainerComponent } from './initial-page/central-container/central-container.component';
import { HeaderComponent } from './initial-page/header/header.component';

import {
  MatButtonModule, MatCardModule, MatDialogModule, MatInputModule, MatTableModule, MatStepperModule,
  MatRadioModule, MatToolbarModule, MatMenuModule, MatIconModule, MatProgressSpinnerModule, MatTabsModule
} from '@angular/material';

@NgModule({
  declarations: [
    AppComponent,
    LoginBoxComponent,
    RegisterBoxComponent,
    CentralContainerComponent,
    HeaderComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatIconModule,
    MatCardModule,
    MatRadioModule,
    MatStepperModule,
    MatTabsModule,
    CommonModule,
    MatToolbarModule,
    MatInputModule,
    MatDialogModule,
    MatTableModule,
    MatMenuModule,
    MatProgressSpinnerModule,
    FormsModule, ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent],
  exports: [
    CommonModule,
    MatToolbarModule,
    MatButtonModule,
    MatCardModule,
    MatInputModule,
    MatDialogModule,
    MatTableModule,
    MatMenuModule,
    MatIconModule,
    MatProgressSpinnerModule
  ],
})
export class AppModule { }
