import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {PlantComponent} from './plant/plant.component';
import {TrayComponent} from './tray/tray.component';
import {GrowConfigComponent} from "./growconfig/growconfig.component";
import {DashboardComponent} from "./dashboard/dashboard.component";
import {TrayConfigComponent} from "./tray-config/tray-config.component";

const routes: Routes = [
    {path: '', redirectTo: '/dashboard', pathMatch: 'full'},
    {path: 'dashboard', component: DashboardComponent},
    {path: 'plant', component: PlantComponent},
    {path: 'growconfig', component: GrowConfigComponent},
    {path: 'trayconfig', component: TrayConfigComponent},
    {path: 'tray', component: TrayComponent},
];

// configures NgModule imports and exports
@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
