import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Plant} from './plant.model';
import {MessageService} from "primeng/api";
import {PlantService} from "./plant.service";

@Component({
    selector: 'app-plant',
    templateUrl: './plant.component.html',
    styleUrls: ['./plant.component.scss']
})
export class PlantComponent implements OnInit {

    plants: Plant[];
    plant: Plant;
    plantDialog: boolean;
    submitted: boolean;

    constructor(private plantService: PlantService, private httpClient: HttpClient, private messageService: MessageService) {
    }


    ngOnInit(): void {
        this.plant = {};
        this.loadData();
    }

    loadData() {
        console.log('loading data...');
        this.plantService.read().subscribe(data => this.plants = data)
    }

    delete(id: number) {
        this.plantService.delete(id).subscribe(() => console.log('deleted'))
        this.loadData();
    }

    edit(plant: Plant) {
        this.plant = {...plant};
        this.plantDialog = true;
    }

    save() {
        this.submitted = true;
        if (this.plant['Id'] === null || this.plant.Id === undefined) {
            this.plantService.create(this.plant)
                .subscribe(value => console.log(JSON.stringify(value)));
            this.messageService.add({
                severity: 'success',
                summary: 'Successful',
                detail: 'GrowConfig created',
                life: 3000
            });
        } else {
            this.plantService.update(this.plant)
                .subscribe(value => console.log(JSON.stringify(value)));
            this.messageService.add({
                severity: 'success',
                summary: 'Successful',
                detail: 'GrowConfig updated',
                life: 3000
            });
        }

        this.loadData();
        this.plantDialog = false;
        this.plant = {};
    }

    openNew() {
        this.plant = {};
        this.submitted = false;
        this.plantDialog = true;
    }

    hideDialog() {
        this.plantDialog = false;
        this.submitted = false;
    }

    getPlantImage(plant: Plant): string {
        return `assets/images/plants/${plant.Name.toLocaleLowerCase()}.jpg`
    }
}
