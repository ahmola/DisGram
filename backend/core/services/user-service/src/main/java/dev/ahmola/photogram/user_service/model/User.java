package dev.ahmola.photogram.user_service.model;

import dev.ahmola.photogram.common_domain.domain.BaseEntity;
import jakarta.persistence.*;
import lombok.*;

@NoArgsConstructor
@Getter@Setter@ToString
@Builder
@Entity
@Table(name = "users")
public class User extends BaseEntity {
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Id
    private Long id;

    @Column(nullable = false, unique = true, length = 50)
    private String username;

    @Column(nullable = false)
    private String passwordHash;

    @Lob
    private String bio;

    @Column(length = 500)
    private String avatarUrl;

}
