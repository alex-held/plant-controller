import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppComponent} from './app.component';
import {HttpClientModule} from '@angular/common/http';
import {AppRoutingModule} from './app-routing.module';
import {TrayComponent} from './tray/tray.component';
import {PlantComponent} from './plant/plant.component';
import {RouterModule} from '@angular/router';

import {TableModule} from 'primeng/table';
import {ToastModule} from 'primeng/toast';
import {CalendarModule} from 'primeng/calendar';
import {DialogModule} from 'primeng/dialog';
import {ButtonModule} from 'primeng/button';
import {DropdownModule} from 'primeng/dropdown';
import {InputTextModule} from 'primeng/inputtext';
import {ToolbarModule} from 'primeng/toolbar';
import {InputNumberModule} from 'primeng/inputnumber';
import {FormsModule} from '@angular/forms';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MessageService} from "primeng/api";
import {GrowConfigComponent} from "./growconfig/growconfig.component";
import {DashboardComponent} from './dashboard/dashboard.component';
import {DataViewModule} from "primeng/dataview";
import {FlexLayoutModule} from "@angular/flex-layout";
import {TrayConfigComponent} from "./tray-config/tray-config.component";
import {ToggleButtonModule} from "primeng/togglebutton";
import {FileUploadModule} from "primeng/fileupload";

@NgModule({
    declarations: [
        AppComponent,
        TrayComponent,
        PlantComponent,
        GrowConfigComponent,
        DashboardComponent,
        TrayConfigComponent
    ],
    imports: [
        ToggleButtonModule,
        FlexLayoutModule,
        DataViewModule,
        BrowserModule,
        BrowserAnimationsModule,
        TableModule,
        CalendarModule,
        DialogModule,
        DropdownModule,
        ButtonModule,
        ToastModule,
        InputTextModule,
        HttpClientModule,
        FileUploadModule,
        ToolbarModule,
        FormsModule,
        InputNumberModule,
        AppRoutingModule,
        RouterModule,
        TableModule,
        ToolbarModule
    ],
    providers: [MessageService],
    bootstrap: [AppComponent]
})
export class AppModule {
}
