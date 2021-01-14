import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {TechComponent} from './tech/tech.component';
import {PlantComponent} from './plant/plant.component';
import {TrayComponent} from './tray/tray.component';

const routes: Routes = [
    {path: 'tech', component: TechComponent},
    {path: 'plant', component: PlantComponent},
    {path: 'tray', component: TrayComponent},
];

// configures NgModule imports and exports
@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
