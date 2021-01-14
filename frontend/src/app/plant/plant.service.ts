import {Injectable} from "@angular/core";
import {Observable} from "rxjs";
import {Plant} from "./plant.model";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";

@Injectable({providedIn: "any"})
export class PlantService {
    private baseUrl = `${environment.apiUrl}/v1/plant/`;

    constructor(private httpClient: HttpClient) {
    }

    create(growConfig: Plant): Observable<Plant> {
        const url = this.baseUrl
        console.log('POST | ' + url);
        return this.httpClient.post<Plant>(url, growConfig);
    }

    read(): Observable<Plant[]> {
        const url = this.baseUrl
        console.log('GET | ' + url);
        return this.httpClient.get<Plant[]>(url);
    }

    getById(id: number): Observable<Plant> {
        const url = `${this.baseUrl}${id}`
        console.log('GET | ' + url);
        return this.httpClient.get<Plant>(url);
    }


    update(growConfig: Plant): Observable<Plant> {
        const url = `${this.baseUrl}${growConfig.Id}`
        console.log('PUT | ' + url);
        return this.httpClient.put<Plant>(url, growConfig);
    }

    delete(id: number): Observable<any> {
        const url = `${this.baseUrl}${id}`
        console.log('DELETE | ' + url);
        return this.httpClient.delete<any>(url);
    }

}
