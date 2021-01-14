import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Plant} from './plant.model';
import {environment} from '../../environments/environment';
import {MessageService} from "primeng/api";

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

    constructor(private httpClient: HttpClient, private messageService: MessageService) {
    }


    ngOnInit(): void {
        this.plant = {};
        this.loadData();
    }

    loadData() {
        const url = environment.apiUrl + '/v1/plant/';
        console.log('loading data...');
        this.httpClient.get<Plant[]>(url)
            .subscribe(value => this.plants = value);
    }

    delete(id: number) {
        const url = environment.apiUrl + '/v1/plant/' + id;
        console.log('DELETE | ' + url);
        this.httpClient.delete(url)
            .subscribe(value => console.log(value));
        this.loadData();
    }

    edit(plant: Plant) {
        this.plant = {...plant};
        this.plantDialog = true;
    }

    save() {
        console.log(JSON.stringify(this.plant));
        console.log(this.plant.Id);

        this.submitted = true;

        let url = environment.apiUrl + '/v1/plant/';
        if (this.plant['Id'] !== null) {
            url = `${url}${this.plant['Id']}`;
            console.log('PUT | ' + url);
            this.httpClient.put<Plant>(url, this.plant)
                .subscribe(value => console.log(JSON.stringify(value)));

            this.messageService.add({severity: 'success', summary: 'Successful', detail: 'Plant Updated', life: 3000});
        } else {
            console.log('POST | ' + url);
            this.httpClient.post<Plant>(url, this.plant)
                .subscribe(value => console.log(JSON.stringify(value)));

            this.messageService.add({severity: 'success', summary: 'Successful', detail: 'Plant Created', life: 3000});
        }

        this.loadData();

        //  this.plants = [...this.plants];
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

}
