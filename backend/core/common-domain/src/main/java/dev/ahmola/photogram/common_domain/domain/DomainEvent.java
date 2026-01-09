package dev.ahmola.photogram.common_domain.domain;

import java.time.LocalDateTime;


public interface DomainEvent {
    LocalDateTime occuredAt();
}
