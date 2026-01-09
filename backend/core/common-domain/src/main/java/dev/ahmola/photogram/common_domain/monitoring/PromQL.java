package dev.ahmola.photogram.common_domain.monitoring;

public class PromQL {

    private PromQL(){
        // private empty constructor
    }

    public static final String REQUEST_RATE_BY_STATUS =
            "sum(rate(http_server_requests_seconds_count[1m])) by (status)";

    public static final String AVERAGE_MEMORY_USAGE =
            "avg(container_memory_usage_bytes{namespace=\"photogram\"})";
}
