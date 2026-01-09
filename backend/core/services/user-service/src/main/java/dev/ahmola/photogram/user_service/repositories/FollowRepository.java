package dev.ahmola.photogram.user_service.repositories;

import dev.ahmola.photogram.user_service.model.Follow;
import dev.ahmola.photogram.user_service.model.FollowId;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface FollowRepository extends JpaRepository<Follow, FollowId> {
}