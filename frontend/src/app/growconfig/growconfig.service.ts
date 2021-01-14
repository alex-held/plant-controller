import {Injectable} from "@angular/core";
import {Observable} from "rxjs";
import {GrowConfig} from "./growconfig.model";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";

@Injectable({providedIn: "any"})
export class GrowConfigService {
    private baseUrl = `${environment.apiUrl}/v1/growconfig/`;

    constructor(private httpClient: HttpClient) {
    }

    create(growConfig: GrowConfig): Observable<GrowConfig> {
        const url = this.baseUrl
        console.log('POST | ' + url);
        return this.httpClient.post<GrowConfig>(url, growConfig);
    }

    read(): Observable<GrowConfig[]> {
        const url = this.baseUrl
        console.log('GET | ' + url);
        return this.httpClient.get<GrowConfig[]>(url);
    }

    getById(id: number): Observable<GrowConfig> {
        const url = `${this.baseUrl}${id}`
        console.log('GET | ' + url);
        return this.httpClient.get<GrowConfig>(url);
    }


    update(growConfig: GrowConfig): Observable<GrowConfig> {
        const url = `${this.baseUrl}${growConfig["Id"]}`
        console.log('PUT | ' + url);
        return this.httpClient.put<GrowConfig>(url, growConfig);
    }

    delete(id: number): Observable<any> {
        const url = `${this.baseUrl}${id}`
        console.log('DELETE | ' + url);
        return this.httpClient.delete<any>(url);
    }

}
