import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppComponent} from './app.component';
import {TechComponent} from './tech/tech.component';
import {HttpClientModule} from '@angular/common/http';
import {AppRoutingModule} from './app-routing.module';
import { TrayComponent } from './tray/tray.component';
import { PlantComponent } from './plant/plant.component';
import {RouterModule} from '@angular/router';

@NgModule({
    declarations: [
        AppComponent,
        TechComponent,
        TrayComponent,
        PlantComponent
    ],
    imports: [
        BrowserModule,
        HttpClientModule,
        AppRoutingModule ,
        RouterModule
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule {
}
