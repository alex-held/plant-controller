import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Plant} from './plant.model';
import {environment} from '../../environments/environment';

@Component({
    selector: 'app-plant',
    templateUrl: './plant.component.html',
    styleUrls: ['./plant.component.scss']
})
export class PlantComponent implements OnInit {

    plants: Plant[] = [];

    constructor(private httpClient: HttpClient) {
    }

    ngOnInit(): void {
        const url = environment.apiUrl + '/v1/plant/';
        console.log(url);
        this.httpClient.get<Plant[]>(url)
            .subscribe(value => this.plants = value);
    }


    add(plantname: string) {
        const url = environment.apiUrl + '/v1/plant/';
        console.log(url);
        this.httpClient.post<Plant>(url, {name: plantname})
            .subscribe(value => console.log(value));
    }
}
