import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {InitScreenComponent} from "./init-screen/init-screen.component";
import {UserScreenComponent} from "./user-screen/user-screen.component";

const routes: Routes = [
  {path: '', component: InitScreenComponent},
  {path: 'user', component: UserScreenComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
