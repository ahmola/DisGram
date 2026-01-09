package dev.ahmola.photogram.common_domain.monitoring;

public final class LokiQL {
    private LokiQL() {
        // private empty constructor
    }

    public static final String POST_SERVICE_ERROR =
            "{app=\"post-service\"} |= \"ERROR\"";

    public static final String USER_SERVICE_KAFKA =
            "{app=\"user-service\"} |= \"KafkaProducer\"";
}
