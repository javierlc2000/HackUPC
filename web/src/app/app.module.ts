import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginBoxComponent } from './initial-page/login-box/login-box.component';
import { RegisterBoxComponent } from './initial-page/register-box/register-box.component';
import { CentralContainerComponent } from './initial-page/central-container/central-container.component';
import { HeaderComponent } from './initial-page/header/header.component';

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
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
